Ensure local ~/.ssh/config file contains an entry for "deploy.target"

Deploy:
```./deploy.sh```

See logs:
```journalctl -u maint.service```

Turn off deployed service:
```sudo systemctl start maint.service```

