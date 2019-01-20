# go-ardupilot

Project to try and get a ardupilot powered drone flying based on the gobot framework


## Prerequisites

### Flight Controller
- Pixhawk 4

### Companion Computer
- Rapsberry Pi 3
- gstreamer
- http://z25.org/static/_rd_/videostreaming_intro_plab/


#### Personal Notes Rapsberry Pi
1. Probably don't need to use upstream repo
2. sudo apt-get install gstreamer1.0-tools \
  gstreamer1.0-plugins-base \
  gstreamer1.0-plugins-good \
  gstreamer1.0-plugins-bad \
  gstreamer1.0-plugins-ugly


Run ```rpi-update``` for v4L2 drivers and ```modprobe bcm2835-v4l2``` to load it if camera doesn't work?


##### Getting Stream to work

Run following command on PI
```
gst-launch-1.0 v4l2src device=/dev/video0 ! video/x-raw,width=640,height=480,framerate=30/1 ! videoconvert ! jpegenc !  rtpjpegpay !  udpsink host=<ip of recieveing host> port=xxxx
```

Recieving Data
```
gst-launch-1.0 udpsrc port=5200 !  application/x-rtp, encoding-name=JPEG,payload=26 !  rtpjpegdepay !  jpegdec ! videoconvert ! autovideosink
```