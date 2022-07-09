from db_lib import *

db, cursor = connect_to_db(DATABASES[1])

tables = get_all_tables(cursor)

for show_name in tables:
    print(f"Starting with {show_name}")
    show = get_show(show_name)
    serial = 1
    for episode in show:
        season_number = episode[2]
        episode_number = episode[1]
        try:
            query = f"UPDATE {show_name} SET serial_number = {serial} \
                  WHERE season_number = {season_number} AND episode_number = {episode_number}"
            print(f"Query: {query}")
            cursor.execute(query)
            serial += 1
        except mysql.connector.Error as e:
            print(f"MySQL Error: {e}")
    print(f"Done with {show_name}")
    db.commit()