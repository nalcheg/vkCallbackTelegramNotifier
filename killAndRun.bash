#!/bin/bash

sudo killall -9 vkCallbackTelegramNotifier
sleep 3
sudo cp ./go_build_main_go /usr/local/bin/vkCallbackTelegramNotifier
vkCallbackTelegramNotifier >> /var/log/vkCallbackTelegramNotifier.log &
