import subprocess

done = True
link = ''
while done:
    link = str(input('Introdu un link PlayList sau link normal : '))
    if link and len(link) >15:
        done = False
    else:
        print('Link gresit incerca dinou ')

# de facut logica la gasirea playlisturi imagini cover de la albumele originale
# de avut optiunea de alege intre ele 

template =' %(playlist)s/%(playlist_index)s - %(title)s.%(ext)s'

command = [
    "yt-dlp",
    "-o", template,
    "--write-thumbnail",
    "-f", "bestaudio",
    "-t","mp3",
    "-P","/home/bivol/Music",
    link
]


subprocess.run(command)

