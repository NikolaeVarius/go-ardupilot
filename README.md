# go-ardupilot

Project to try and get a ardupilot powered drone flying based on the gobot framework


## Prerequisites

### Flight Controller
- Pixhawk 4

### Companion Computer
- Rapsberry Pi 3
- gstreamer
- http://z25.org/static/_rd_/videostreaming_intro_plab/
- https://github.com/thaytan/gst-rpicamsrc
- http://ardupilot.org/dev/docs/raspberry-pi-via-mavlink.html

### Installing on Arch

#### Personal Notes Rapsberry Pi
1. Probably don't need to use upstream repo
2. sudo apt-get install gstreamer1.0-tools \
  gstreamer1.0-plugins-base \
  gstreamer1.0-plugins-good \
  gstreamer1.0-plugins-bad \
  gstreamer1.0-plugins-ugly


Run ```rpi-update``` for v4L2 drivers and ```modprobe bcm2835-v4l2``` to load it if camera doesn't work?


##### Getting Stream to work
Starting Recieve Server First

Recieving Data
```
gst-launch-1.0 udpsrc port=5000 \
! gdpdepay \
! rtph264depay \
! avdec_h264 \
! videoconvert \
! autovideosink sync=false
```


Run following command on PI
Using v4l2src
```
gst-launch-1.0 v4l2src device=/dev/video0 \
! video/x-h264,width=640,height=480,framerate=30/1 \
! h264parse \
! queue \
! rtph264pay config-interval=1 pt=96 \
! gdppay \
! udpsink host=<ip of recieveing host> port=5000
```

Using rpicamsrc
```
gst-launch-1.0 rpicamsrc bitrate=1000000 \
! 'video/x-h264,width=640,height=480' \
! h264parse \
! queue \
! rtph264pay config-interval=1 pt=96 \
! gdppay \
! udpsink host=[INSERT_IP_ADDR] port=5000
```

## Credits
Using classifier data from https://github.com/hybridgroup/gocv/blob/master/data/haarcascade_frontalface_default.xml

## Useful Links
### Mavlink Docs
https://mavlink.io/en/

### Mavlink for Golang

https://github.com/gswly/gomavlib 
https://github.com/liamstask/go-mavlink (Dead?)
https://github.com/ungerik/go-mavlink (Dead?)
https://github.com/SpaceLeap/go-mavlink (Dead?)
