FROM golang:1.21-buster

RUN apt-get update
RUN apt-get install -y git
RUN apt-get update && apt-get install -y -q --no-install-recommends \
    build-essential \
    libfontconfig1-dev \
    libfreetype6-dev \
    $jpeg \
    libpng-dev \
    $ssl \
    libx11-dev \
    libxext-dev \
    libxrender-dev \
    python \
    zlib1g-dev \
    xfonts-75dpi \
    xfonts-base \
    libjpeg62-turbo \
    fontconfig \ 
    && rm -rf /var/lib/apt/lists/*

RUN wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.5/wkhtmltox_0.12.5-1.buster_amd64.deb
RUN dpkg -i wkhtmltox_0.12.5-1.buster_amd64.deb
RUN apt --fix-broken install

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./app/main.go

EXPOSE 3000

CMD ["./out/app"]