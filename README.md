# Nyabi-Go

Go 표준 라이브러리 기반의 이미지 배치 변환 도구입니다.
지정된 경로 내의 이미지 파일을 탐색하여 JPEG 포맷으로 일괄 변환합니다.

## Features
- PNG/JPG/JPEG 파일 스캔 및 확장자 필터링
- image 패키지를 이용한 이미지 디코딩 및 재인코딩
- 출력 디렉토리 생성 및 파일 입출력 처리
- 순차적 처리를 통한 자원 관리

## Prerequisites
Go 1.26

**macOS (Homebrew)**
```zsh
brew install go
```

**Arch Linux**
```zsh
sudo pacman -S go
```

## Build
```zsh
go build -o nyabi-img main.go
```

## Run
```zsh
./nyabi-img
```