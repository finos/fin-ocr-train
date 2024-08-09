"""
Copyright (c) 2024 Discover Financial Services
"""
import os, sys, random, pathlib, subprocess, time

micrCharacters = 'ABCD0123456789 '
font = 'micr'
modelName = 'micr_FT'

if len(sys.argv) != 3:
    print("Usage: gen_training_data <count> <output-dir>")
    sys.exit(1)

count = int(sys.argv[1])
output_dir = sys.argv[2]

pathlib.Path(output_dir).mkdir(parents=True, exist_ok=True)

possibleExpos = [0, 1, 2, 3]

for i in range(1,count+1):
    micrText = ''.join(random.choices(micrCharacters, k=40))
    base = f'{modelName}_{i}'
    gtFile = os.path.join(output_dir, f'{base}.gt.txt')
    with open(gtFile, 'w') as output_file:
        output_file.writelines([micrText])
    
    exposure = possibleExpos[i % len(possibleExpos)]

    if i % 100 == 0:
        print("\n----  Iteration Number:", i, " ---- ", time.ctime()," ----\n")
    
    subprocess.run([
        'python3',
        '../TextRecognitionDataGenerator/trdg/run.py',
        # Count
        '-c=1',
        # Distortion
        '-d=3',
        # Skew
        '-k=2',
        '-rk',
        # Random Sequence
        #'-rs',
        # Background
        '-b=1',
        # Blur
        '-bl=2',
        '-rbl',
        # Height
        '-f=64',
        # Case
        '-ca=upper',
        # Font
        '-ft=/home/ubuntu/.fonts/micr.ttf',
        # Input File
        f'-i={gtFile}',
        # Output Boxes
        '-obb=2',
        # File Extension
        '-e=tif',
        # Stroke Width
        f'-stw={exposure}',
        # Output Directory
        f'--output_dir={output_dir}',
    ])
    
    # Rename the output files to match the groundtruth file
    os.rename(os.path.join(output_dir, micrText+'_0.tif'), os.path.join(output_dir, base+'.tif'))
    os.rename(os.path.join(output_dir, micrText+'_0.box'), os.path.join(output_dir, base+'.box'))


