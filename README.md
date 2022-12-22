# 최종 KPI : NFT 현금거래 지원을 가능하도록 구성하시오. (~12/23)
---
## 1차 개발 KPI : 최소한의 기능으로 동작 가능하도록 구성하시오.
  ### 🔖 필요 개발 기능
  - NFT Item Schema
  - 컬렉션 생성, 조회 기능
  - 그룹 생성, 조회 기능
  - NFT Item 생성 기능
  - NFT Item 전체 조회 기능
  - NFT Item 페이징 조회 기능 *(1 ~ 10개씩, Page & Limit를 주어 처리)*
  - NFT Item 단일 조회 기능 (By ID)
  - NFT Item FK를 통한 JOIN 조회 API

  - 유저 JWT 회원가입
  - 유저 JWT 로그인
  - 유저 JWT 로그아웃
  - 경로(라우트)별로 접근 제한(권한) 미들웨어
  - 결제 pg API

## 2차 개발 KPI : 요청에 맞추어 다듬어 완성하고 안정적으로 동작할 수 있도록하시오.
---

### 프로젝트 구조 설명 (MVC Pattern)
 ┣ 📂configs -> ``` DB 연결이 현재 포함되어있음.```   
 ┃ ┗ 📜db.go    
 ┣ 📂controller -> ``` API 핸들러 작성 ```   
 ┃ ┣ 📜item_controller.go    
 ┃ ┗ 📜user_controller.go    
 ┣ 📂model  -> ``` DB Schema 작성 ```  
 ┃ ┣ 📜item_schema.go  -> ``` Item관련 ```  
 ┃ ┗ 📜user_schema.go  -> ``` User관련 ```  
 ┣ 📂routes  -> ``` Router 설정 폴더 ```  
 ┃ ┗ 📜router.go    
 ┣ 📜go.mod  -> ``` 모듈&패키지 ```   
 ┣ 📜go.sum     
 ┗ 📜main.go  -> ``` Root ```   
