FROM alios32/sql-migrate

COPY migrations/trade-bot /migrations/trade-bot

WORKDIR /

ENTRYPOINT ["migrations/trade-bot/migrator.sh"]
