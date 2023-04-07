import urllib.request
import json

def get_all_video_in_channel(channel_id):
    api_key = "AIzaSyDXCuguEKvISldv2uVWXG0itvKRFzlbueU"

    base_video_url = 'https://www.youtube.com/watch?v='
    base_search_url = 'https://www.googleapis.com/youtube/v3/search?'

    first_url = base_search_url + f'key={api_key}&channelId={channel_id}&part=snippet,id&order=date&maxResults=25'

    video_links = []
    url = first_url
    while True:
        with urllib.request.urlopen(url) as inp:
            resp = json.load(inp)

        for i in resp['items']:
            if i['id']['kind'] == "youtube#video":
                video_links.append(base_video_url + i['id']['videoId'])

        try:
            next_page_token = resp['nextPageToken']
            url = first_url + f'&pageToken={next_page_token}'
        except:
            break
    return video_links

#     UtxtaUMNSvxq_Xg","externalId":"UCfeMEuhdUtxtaUMNSvxq_Xg","keywords":"fryingpa
_ = [print(x) for x in get_all_video_in_channel("UCfeMEuhdUtxtaUMNSvxq_Xg")]