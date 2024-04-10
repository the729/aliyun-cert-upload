# aliyun-cert-upload
Upload a set of TLS cert/key to Aliyun Yundun DCAS.

A handy tool to work with acme / let's encrypt auto renewal.

## Usage

```
export Ali_Key=xxx
export Ali_Secret=xxx
aliyun-cert-upload -base-name mycert -cert /path/to/fullchain.pem -key /path/to/key.pem
```

If used with acme.sh reload command, you probably want to set Ali AK/SK env in ~/.bashrc or ~/.acme.sh/acme.sh.env.

```
acme.sh -i -d xxx.example.com \
  --fullchain-file /path/to/fullchain.pem \
  --key-file /path/to/key.pem \
  --reloadcmd "aliyun-cert-upload -base-name mycert -cert /path/to/fullchain.pem -key /path/to/key.pem"
```
