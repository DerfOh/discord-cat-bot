FROM scratch
ADD ca-bundle.crt /etc/ssl/certs/
ADD discord-cat-bot /
ADD config.json /
COPY commands/sounds /commands/sounds
COPY notifications/sounds /notifications/sounds
CMD ["/discord-cat-bot"]