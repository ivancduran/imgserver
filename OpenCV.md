Install opencv in CentOS
<<<<<<< HEAD
Opencv can be install in CentOS in two ways:
=======
Opencv can be install in CentOS:
>>>>>>> 67cdfc97ac817b8db8f96a8c0005e749b7dfae4c

1) Install from available yum repository.
But till the time of writing this post opencv-2.0.0 repository are available. If you need latest version of opencv then go for 2. option.
2) Opencv installation from source.

1) Install opencv from yum repo
Before installation of opencv-python from yum repository we have to install require library numpy.

$ sudo yum install python-devel python-nose python-setuptools gcc gcc-gfortran gcc-c++ blas-devel lapack-devel atlas-devel
$ sudo easy_install pip
$ sudo pip install numpy==1.6.1
If you haven’t installed opencv rpm then you can download from here and install it first . After installation of rpm now opencv will be available for installation.


yum install numpy opencv*

sudo cp facedetect /usr/local/bin

The problem here is that the package for Fedora doesn't just have a different name than what's mentioned in the facedetect documentation — it also puts its files in /usr/share/OpenCV/ instead of /usr/share/opencv, which is where facedetect expects them to be. This is annoying but very easy to fix. If you have the same problem, open the facedetect script with a text editor and replace opencv with OpenCV in this line:

 `` 'HAAR_FRONTALFACE_ALT2': '/usr/share/opencv/haarcascades/haarcascade_frontalface_alt2.xml'``

