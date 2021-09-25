FROM challisa/easyocr

ADD ./easyocr-docker easyocr-docker
ADD ./ocr.py ocr.py
ADD ./example example

RUN python ocr.py example/fleet.png

EXPOSE 8080
ENTRYPOINT ["./easyocr-docker"]