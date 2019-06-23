FROM scratch
ADD ca-bundle.crt /etc/ssl/certs/
ADD discord-cat-bot /
ADD config.json /
COPY commands/sounds /commands/sounds
CMD ["/discord-cat-bot"]