FROM scratch
ADD ca-bundle.crt /etc/ssl/certs/
ADD discord-cat-bot /
ADD config.json /
CMD ["/discord-cat-bot"]