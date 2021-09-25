import easyocr
import sys

reader = easyocr.Reader(['en', 'ja'], gpu=False)

print(reader.readtext(sys.argv[1], detail=0)[0])
