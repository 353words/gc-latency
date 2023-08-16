import json
import pandas as pd


def load_df(file_name):
    with open(file_name) as fp:
        return pd.json_normalize(json.loads(line) for line in fp)


map_df = load_df('out/map.json')
slice_df = load_df('out/slice.json')
str_df = load_df('out/str.json')

for col in map_df.columns:
    if not col.startswith('cpu'):
        continue
    print(col)
    print('\tmap  : {:12.3f}'.format(map_df[col].mean()))
    print('\tslice: {:12.3f}'.format(slice_df[col].mean()))
    print('\tstr  : {:12.3f}'.format(str_df[col].mean()))
