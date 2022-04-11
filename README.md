# Online PDF Editor 🥳

website that can edit PDF's

[![Spellcheck CI](https://github.com/dipankardas011/PDF-Editor/actions/workflows/spellcheck.yml/badge.svg)](https://github.com/dipankardas011/PDF-Editor/actions/workflows/spellcheck.yml) [![pages-build-deployment](https://github.com/dipankardas011/PDF-Editor/actions/workflows/pages/pages-build-deployment/badge.svg)](https://github.com/dipankardas011/PDF-Editor/actions/workflows/pages/pages-build-deployment)

## Kubernetes repo for this
[`IAC repo`](https://github.com/dipankardas011/PDF-Editor-IAC)

### Tech Stack
* GO
* Docker
* HTML
<!--  redis DB -->


## Current Deployment is on Heroku

# Website
![](./coverpage.png)

# Website Link
[Click Here](https://pdf-editor-tool.herokuapp.com/)

## WORK 🚧
Work | Status
-|-
Backend | ✅
Database | 🚧


# Flow of the program using Graphs` OUTDATED❕`
```mermaid
flowchart LR;
    XX[START]:::white-->web{Website};
    web{Website}-->B{file1 uploaded};
    web{Website}-->C{file2 uploaded};
    DD{Download Link}-->web{Website};

    classDef green color:#022e1f,fill:#00f500;
    classDef red color:#022e1f,fill:#f11111;
    classDef white color:#022e1f,fill:#fff;
    classDef black color:#fff,fill:#000;

    B--upload 1-->S[GO Server]:::green;
    C--upload 2-->S[GO Server]:::green;

    S[GO server]-->DD{Download Link}

```

# How to Run

```bash
./Runner.sh
```

## connect to the redis db `UNDER DEVELOPMENT`

```bash
docker ps
docker exec it <container id> bash
redis-cli
```

## connect to the frontend

```url
localhost:80
```

Happy Coding 🥳
