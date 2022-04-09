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


# Website
![](./coverpage.png)

# Website Link
[Click Here](https://pdf-editor-tool.azurewebsites.net)

## WORK 🚧
Work | Status
-|-
Backend | ✅
Database | 🚧


### Links

[GO REDIS](https://github.com/gomodule/redigo)

[PDF search for go](https://pkg.go.dev/search?q=pdf)

![](./techStack.svg)

# Flow of the program using Graphs` OUTDATED❕`
```mermaid
flowchart LR;
    XX[START]-->web[website];
    web[website]-->B{file1 uploaded};
    web[website]-->C{file2 uploaded};
    web[website]-->DD{Download Link};

    classDef green color:#022e1f,fill:#00f500;
    classDef red color:#022e1f,fill:#f11111;
    classDef white color:#022e1f,fill:#fff;
    classDef black color:#fff,fill:#000;

    B--success_upload-->S[Submit button]:::green;
    B--oerror-->E[Submit button]:::red;
    C--success_upload-->S[Submit button]:::green;
    C--oerror-->E[DEAD STATE]:::red;

    S--API POST-->backend{processing};
    backend--save/store-->DBA[REDIS]:::white;
    DBA--editing-->src[GOLANG]:::green;
    src--savingResult-->DBA;

    DBA--creating Public Link-->backend;
    DD--GET File using exposed port-->backend:::white;
```

Going by Type1

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