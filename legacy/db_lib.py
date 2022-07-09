import random
import mysql.connector

DATABASES = [
    'unwatched_tv_shows',
    'watched_tv_shows'
]

def connect_to_db(db_name):
    """ Tries to connect to a database, returns tuple (database obj, connection) or None """
    try:
        db = mysql.connector.connect(
        host='localhost',
        user='root',
        password='Not@GoodPassword1234',
        database=f'{db_name}'
        )
    except mysql.connector.Error as e:
        print(f"Something went wrong: {e}")
        return None
    return (db, db.cursor())

def get_all_tables(cursor):
    query = "SHOW TABLES"
    cursor.execute(query)
    tables = []
    for row in cursor:
        table_name = row[0]
        tables.append(table_name)
    return tables

def get_all_shows(db_name):
    try:
        _, cursor = connect_to_db(db_name)
    except ValueError as e:
        return None
    shows = []
    for table in get_all_tables(cursor):
        if table.startswith('_'):
            continue
        shows.append(table)
    return shows

def get_random_season(db_name, show_title):
    try:
        _, cursor = connect_to_db(db_name)
    except ValueError as e:
        return None
    query = f"SELECT DISTINCT season_number FROM {show_title}"
    cursor.execute(query)
    result = cursor.fetchall()
    print(f"Found {len(result)} seasons of {show_title}")
    return random.randint(1, len(result))

def _get_random_episode(db_name, show_title, season):
    try:
        _, cursor = connect_to_db(db_name)
    except ValueError as e:
        return None
    query = f"SELECT DISTINCT title, episode_number FROM {show_title} WHERE season_number = {season}"
    cursor.execute(query)
    result = cursor.fetchall()
    print(f"Found {len(result)} episodes for {show_title} season {season}")
    return random.choice(result)

    
def import_tv_show(db, cursor, show_title, episodes, marked=False):
    # Create new table, assuming filename is the name of the show
    query = f"CREATE TABLE IF NOT EXISTS {show_title}(serial_number INT, episode_number INT, season_number INT, title VARCHAR(256), watched BOOLEAN)"
    cursor.execute(query)
    # Add episodes to the table
    imported = 0
    for episode in episodes:
        print(episode)
        try:
            serial_number = int(episode['number'])
            episode_number = int(episode['episode'])
            season_number = int(episode['season'])
            title = episode['title'].replace('"', '').replace("'", "")
            query = f'INSERT INTO {show_title} VALUES ({serial_number}, {episode_number}, {season_number}, "{title}", {marked})'
            cursor.execute(query)
            imported += 1 
        except ValueError as e:
            print(f"Skipping row {episode} due to {e}")
        except mysql.connector.Error as e:
            print(f"Skipping row {episode} due to mysql error: {e}")

    db.commit()
    print(f"Successfully imported {imported} rows")

def check_show_completion(show_title, cursor):
    query = f"SELECT * FROM {show_title} WHERE watched = 0"
    cursor.execute(query)
    results = cursor.fetchall()
    if len(results) == 0:
        # Move the show to watched_tv_shows
        query = f"ALTER TABLE unwatched_tv_shows.{show_title} RENAME watched_tv_shows.{show_title}"
        cursor.execute(query)
        return True
    # TODO this could be modified to return how many episodes left
    # And that could help with notifying when you're close to finishing a show
    return False

def batch_toggle_episodes(show_title, episodes):
    db, cursor = connect_to_db('unwatched_tv_shows')
    if not show_title in get_all_tables(cursor):
        raise ValueError("Show not found")
    for episode in episodes:
        for v in episode.values():
            if not isinstance(v, int):
                continue
        try:
            query = f"UPDATE {show_title} SET watched = {episode['toggle']} WHERE season_number = {episode['season']} AND episode_number = {episode['episode']}"
            cursor.execute(query)
        except (mysql.connector.Error, KeyError) as e:
            print(f"Skipping episode: {episode} because {e}")
    db.commit()
    return check_show_completion(show_title, cursor)

def get_next(show_title, cursor=None):
    if not cursor:
        _, cursor = connect_to_db('unwatched_tv_shows')
    if not show_title in get_all_tables(cursor):
        raise ValueError('Show not found')
    result = None
    try:
        query = f"SELECT * FROM {show_title} WHERE watched = 0 LIMIT 1"
        cursor.execute(query)
        result = cursor.fetchall()
    except mysql.connector.Error as e:
        print(f"Mysql Error: {e}")
    return result
    
def set_watched_next(show_title):
    db, cursor = connect_to_db('unwatched_tv_shows')
    result = None
    try:
        # First, get the row we're about the change
        episode = get_next(show_title, cursor)
        # Then, update the row
        query = f"UPDATE {show_title} SET watched = 1 WHERE watched = 0 LIMIT 1"        
        cursor.execute(query)
        db.commit()
    except mysql.connector.Error as e:
        print(f"Mysql error: {e}")
    if episode and len(episode) > 0:
        result = get_next(show_title, cursor)
    return result

def get_episode(show_title, season, episode, cursor=None, serial=None):
    # Need a table resolver for situations like this
    if not cursor:
        _, cursor = connect_to_db('unwatched_tv_shows')
    if not show_title in get_all_tables(cursor):
        raise ValueError('Show not found')
    result = None
    try:
        query = f'SELECT * FROM {show_title} WHERE season_number = {season} AND episode_number = {episode}'
        if serial: 
            query = f'SELECT * FROM {show_title} WHERE serial_number = {serial}'
        cursor.execute(query)
        result = cursor.fetchall()
    except mysql.connector.Error as e:
        print(f"Mysql Error: {e}")
    return result

def get_season(show_title, season):
    _, cursor = connect_to_db('unwatched_tv_shows')
    if not show_title in get_all_tables(cursor):
        raise ValueError('Show not found')
    result = None
    try:
        query = f"SELECT * FROM {show_title} WHERE season_number = {season}"
        cursor.execute(query)
        result = cursor.fetchall()
    except mysql.connector.Error as e:
        print(f'Mysql Error: {e}')
    return result

def get_show(show_title):
    _, cursor = connect_to_db('watched_tv_shows')
    if not show_title in get_all_tables(cursor):
        raise ValueError('Show not found')
    result = None
    try:
        query = f"SELECT * FROM {show_title}"
        cursor.execute(query)
        result = cursor.fetchall()
    except mysql.connector.Error as e:
        print("MySQL Error: {e}")
    return result

def toggle_single_episode(show_title, season, episode_number):
    db, cursor = connect_to_db('unwatched_tv_shows')
    episode = get_episode(show_title, season, episode_number, cursor=cursor)
    if not episode:
        return None
    try:
        watched = int(episode[0][-1])
    except (TypeError, IndexError) as e:
        print(f"{e}: {episode}")
        return None
    try:
        query = f"UPDATE {show_title} SET watched = {abs(watched - 1)} WHERE season_number = {season} AND episode_number = {episode_number}"
        cursor.execute(query)
        db.commit()
    except mysql.connector.Error as e:
        print(f"MySQL Error: {e}")
    return get_episode(show_title, season, episode_number, cursor=cursor)
