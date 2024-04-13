FROM alpine

ADD ./bin/habit-tracking /root/habit-tracking

WORKDIR /root

RUN pwd
RUN ls -la /root

ENTRYPOINT ["/bin/sh","-c","/root/habit-tracking"]