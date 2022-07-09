import sys, csv, argparse, os.path, pathlib

from db_lib import *

UNWATCHED_TV_SHOWS = 'unwatched_tv_shows'
WATCHED_TV_SHOWS = 'watched_tv_shows'

marked = False
database = UNWATCHED_TV_SHOWS

def get_episodes_from_csv(filename):
    print(filename)
    episodes = []
    try:
        with open(filename, newline='') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                episodes.append(row)
    except FileNotFoundError:
        print(f"'{filename}' : File not found")
        sys.exit(1)
    return episodes

parser = argparse.ArgumentParser(description='Import a CSV tv show')
parser.add_argument('file', metavar='/path/to/csv', help='The csv file to import')
parser.add_argument('--watched', action='store_true')
args = parser.parse_args()


    
filename = args.file
if args.watched:
    marked = True
    database = WATCHED_TV_SHOWS

episodes = []
episodes = get_episodes_from_csv(filename)

# TODO change show title to a required arg 
# Once that's done, change the table schema for proper name storage
show_title = os.path.basename(filename)[:-4]
    
if len(episodes) == 0:
    print("No episodes found")
    sys.exit(0)    

try:
    db, cursor = connect_to_db(database)
except ValueError as e:
    raise SystemExit(f'Error connecting to db {database}')

import_tv_show(db, cursor, show_title.lower(), episodes, marked=marked)


'''
|=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-|
| Deprecated almost immediately, but kept for posterity's sake             |
| I wanted to see if I could do it, and it took me longer than I expected  |
| Fun way to spend a Friday night though                                   | 
|     *didn't take me all night                                            |
|-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=|

def normalize_episode_numbers(episodes):
    for i,v in enumerate(episodes):
        if int(v['season']) > int(episodes[i-1]['season']):
            normalized = 0
            while episodes[i+normalized]['season'] == v['season']:
                episodes[i+normalized]['number'] = str(normalized + 1)
                normalized += 1
                if i + normalized == len(episodes) - 1:
                    break
'''