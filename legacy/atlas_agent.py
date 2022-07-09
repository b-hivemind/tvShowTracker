import json, random

from flask import Flask, request

import db_lib

app = Flask(__name__)

def _tv_show_to_json(raw_episode_rows):
    episodes = []
    for row in raw_episode_rows:
        try:
            episodes.append({"serial_number": row[0], "episode_number": row[1], "season_number": row[2], "show_title": row[3], "watched": row[4]})
        except IndexError:
            print(f"Skipping conversion of episode {row}")
    return "{}".format(episodes)



@app.route('/')
def index():
    return "Success: tv_show_tracker_api"

@app.route('/all', methods=['GET', 'POST'])
def get_shows():
    db_name = "unwatched_tv_shows"
    if request.method == "POST":
        try:
            db_name = request.get_json()['type']
        except KeyError:
            return "400 Missing Args\n"
        if db_name not in ('watched_tv_shows', 'unwatched_tv_shows'):
            return "404 Type not found\n"
    all_shows = db_lib.get_all_shows(db_name)
    if all_shows == None:
        return "500: Sorry, please try later\n"
    return "{}\n".format(json.dumps(all_shows))

@app.route('/<show_title>')
def return_show(show_title):
    episodes = []
    try:
        episodes = db_lib.get_show(show_title)
        if episodes == None:
            return "404: No episodes found\n"
    except ValueError:
        return "404: Show not found\n"
    return _tv_show_to_json(episodes)

@app.route('/<show_title>/season/<season>')
def return_season(show_title, season):
    episodes = []
    try:
        episodes = db_lib.get_season(show_title, season)
        if episodes == None:
            return "404: No episodes found\n"
    except ValueError as e:
        return "404: Show not found\n"
    return _tv_show_to_json(episodes)

@app.route('/<show_title>/season/<season>/episode/<episode_number>', methods=['GET', 'POST'])
def return_episode(show_title, season, episode_number):
    episode = []
    try:
        if request.method == 'GET':
            episode = db_lib.get_episode(show_title, season, episode_number)
        else:
            episode = db_lib.toggle_single_episode(show_title, season, episode_number)
        if episode == None:
            return "404: Episode not found\n"
    except ValueError:
        return "404: Show not found\n"
    return _tv_show_to_json(episode)

@app.route('/random', methods=['GET', 'POST'])
def get_random_episode():
    if request.method == 'GET':
        db_name = "watched_tv_shows"
        show_title = random.choice(db_lib.get_all_shows(db_name))
        season = db_lib.get_random_season(db_name, show_title)
    else:
        request_data = request.get_json()
        db_name = request_data.get("type", "watched_tv_shows")
        show_title = request_data.get("show_title", random.choice(db_lib.get_all_shows(db_name)))
        season = request_data.get("season", db_lib.get_random_season(db_name, show_title))
    title, num = db_lib._get_random_episode(db_name, show_title, season)
    return "{}\n".format(json.dumps({'show': show_title, 'season': season, 'episode': {'title': title, 'number': num}}))

@app.route('/<show_title>/next', methods=['POST', 'GET']) 
def set_next(show_title):
    try:
        if request.method == 'POST':
            updated_episode = db_lib.set_watched_next(show_title)
            if  updated_episode == None:
                return "500: Internal Server Error\n"
            return "{}\n".format(json.dumps(_tv_show_to_json(updated_episode)))
        else:
            next_episode = db_lib.get_next(show_title)
            if next_episode == None:
                return "500: Internal Server Error\n"
            return json.dumps(_tv_show_to_json(next_episode))
    except ValueError:
        return f"404: Show {show_title} not found\n"

@app.route('/toggle', methods=['POST'])
def toggle_episodes():
    try:
        request_data = request.get_json()
        show_title = request_data['show_title']
        episodes = request_data['episodes']
    except KeyError:
        return "400: Missing args\n"
    if len(episodes) == 0:
        return "400: No episodes found\n"
    try:
        didfinish = db_lib.batch_toggle_episodes(show_title, episodes)
        return json.dumps({"status": 200, "series_completed":didfinish})
    except ValueError as e:
        if "Show not found" in e.message:
            return f"404: {e}\n"
        else:
            return f"500: Sorry, please try later\n"

if __name__== "__main__":
    app.run(host='0.0.0.0', port=5001, debug=True)
