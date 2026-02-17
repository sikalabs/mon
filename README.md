[SikaLabs (sikalabs.com)](https://sikalabs.com) | [Ondrej Sika (sika.io)](https://sika.io)

# mon

What is **mon**? It is a simple monitoring tool from [SikaLabs](https://github.com/sikalabs). It is inspired by [ondrejsika/mon](https://github.com/ondrejsika/mon), my 10+ years old implementation of simple monitoring tool (written in Python).

## Install mon

Using [slu](https://github.com/sikalabs/slu)

```
slu install-bin mon
```

Using brew

```
brew install sikalabs/tap/mon
```

## Configuration

Create `./mon.yaml` or `/etc/mon/mon.yaml`.

Check out the config file example [mon.example.yaml](./mon.example.yaml)

## Install systemd mon.service

```
mon install
```

Test by

```
systemctl status mon.service
```
