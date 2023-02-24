### 直接调用本机的ffmpeg进行视频截图
[ffmpeg官网](http://ffmpeg.org/)
#### 本机直接使用ffmpeg
```shell
ffmpeg -i ../../../../static/videos/1.mp4 -y -f image2 -t 0.001 ../../../../static/images/1.jpg
```

`-i` 视频路径

`-t` 截图时间

最后的路径是输出文件路径

#### 通过Go调用ffmpeg
使用ffmpeg中的Thumbnail方法即可

#### 参考
https://github.com/langwan/langgo/tree/main/components/ffmpeg