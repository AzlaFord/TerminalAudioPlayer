import subprocess

done = True
link = ''
while done:
    link = str(input('Introdu un link PlayList sau link normal : '))
    if link and len(link) >15:
        done = False
    else:
        print('Link gresit incerca dinou ')

template = str(input('Introdu un nume Ditrectoriu (sau lasa gol ca default) : '))
if len(template) == 0:
    template =' %(playlist)s/%(playlist_index)s - %(title)s.%(ext)s'

command = [
    "yt-dlp",
    "-o", template,
    "--write-thumbnail",
    "-f", "bestaudio",
    "-P","/home/bivol/Music",
    link
]


subprocess.run(command)

