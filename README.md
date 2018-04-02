
Parameters:
==

* i = Image url, ej: i=http://image.com/image.jpg
* w = Width, ej: w=800
* h = Height, ej: h=600
* k = Key, user bucket or folder, ej: k=ivancduran
* t = Transformation, ej: resize, fill, fit, face
* f = Face number, ej: f=1, f=2
* q = Quality, ej: 50
* c = Webp for chrome and opera, ej: c=false
* p = Percent of reduction, ej: p=50

Next features:
==
* Graphical Interface for web written in Angular 2 as a SPA.
* Analysis at delivery logs from Akamai
* Easy integration with CDN: Akamai
* Metrics by IP and Geolocation
* AWS S3 integration
* Bash loader for local storage
* JWT Headers for spesific pictures
* Gif optimization
	* Lossy and Lossless
* Integration with Caddy for http2 (evaluation)
* Letsencrypt (https) for production
* Engines selector and configuration file (jpegtran, jpegrescan, mozjpeg1, mozjpeg2, libjpeg-turbo)
* Gzip only for files lower than 10k
* Nudity detection api

Requerimients:
==

* optinpng
* jpegoptim
* cwebp
* OpenCV


Install optimpng and jpegoptim on Centos 7:
==

yum install optipng
yum install jpegoptim


Install OpenCV on Centos 7:
==

sudo yum install python-devel python-nose python-setuptools gcc gcc-gfortran gcc-c++ blas-devel lapack-devel atlas-devel
sudo easy_install pip
sudo pip install numpy==1.6.1
yum install numpy opencv*
sudo cp facedetect /usr/local/bin

Install webp on Centos 7
==

yum install libjpeg-devel libpng-devel libtiff-devel libgif-devel

wget libwebp-0.5.1.tar.gz
tar xvzf libwebp-0.5.1.tar.gz

cd libwebp-0.5.1
./configure
make
sudo make install


Install webp on windows
==

Download the library and put in that path

C:\libwebp\bin


Test urls:
==

http://localhost:8090/v1/install
http://localhost:8090/v1/up?k=ivan&i=https://golang.org/doc/gopher/fiveyears.jpg
http://localhost:8090/v1/get?k=ivan&i=IJ5XDyD8tCpz6PvE.jpg&t=fit&w=400&h=150


Test Images:
==

https://golang.org/doc/gopher/fiveyears.jpg
https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/Gogophercolor.png/1024px-Gogophercolor.png


docker run -p 27017:27017 --name some-mongo -d mongo
docker run --name easycast-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -d mysql