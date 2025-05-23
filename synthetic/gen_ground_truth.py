"""
Copyright (c) 2024 Capital One
"""
import os, sys, random, pathlib, subprocess

#
# Randomly generate ground truth characters and images (using text2image).
# These pairs of images and ground truth characters are used to train tesseract.
#
# Configuration parameters:
# font - the name of the font
# fontCharacters - the characters of the font
#
font = 'micr'
fontCharacters = 'ABCD0123456789 '

if len(sys.argv) != 3:
    print("Usage: gen_ground_truth <count> <output-dir>")
    sys.exit(1)

count = int(sys.argv[1])
output_dir = sys.argv[2]

pathlib.Path(output_dir).mkdir(parents=True, exist_ok=True)

for i in range(1,count+1):
    micrText = ''.join(random.choices(fontCharacters, k=40))
    base = f'{font}_{i}'
    gtFile = os.path.join(output_dir, f'{base}.gt.txt')
    with open(gtFile, 'w') as output_file:
        output_file.writelines([micrText])

    subprocess.run([
        'text2image',
        f'--font=GnuMicr Thin',
        f'--fonts_dir=./fonts',
        f'--text={gtFile}',
        f'--outputbase={output_dir}/{base}',
        '--max_pages=1',
        '--strip_unrenderable_words',
        '--leading=32',
        '--xsize=3600',
        '--ysize=480',
        '--char_spacing=0.5',
        '--exposure=0',
        f'--unicharset_file=unicharset'
    ])
