TODO from perplexity.ai

프로젝트를 시작할 때는 **계획**을 세우고, 필요한 **하드웨어 및 소프트웨어 준비**를 차근차근 진행하는 것이 중요합니다. 아래는 **Next.js 프론트엔드**와 **Golang 백엔드**를 사용하는 IoT 디스플레이 프로젝트를 시작하기 위한 단계별 가이드입니다.

## **1. 프로젝트 계획 및 요구 사항 정의**

### **1.1 기능 정의**
- **날씨 정보 표시**: 기상청 API에서 날씨 데이터를 받아와 디스플레이에 출력.
- **디스플레이 업데이트 주기**: 일정 시간마다(예: 5분) 날씨 정보를 새로고침.
- **UI 구성**: Next.js로 간단한 웹 페이지를 만들어 날씨 정보를 시각적으로 표시.
- **백엔드 서비스**: Golang으로 기상청 API 데이터를 받아 프론트엔드에 제공.

### **1.2 하드웨어 선택**
- 라즈베리파이 또는 라테판다 중 하나를 선택하고, 필요한 부품을 구매합니다.
  - 예: 라즈베리파이 4, microSD 카드, HDMI 디스플레이 등.

### **1.3 소프트웨어 스택 결정**
- **백엔드**: Golang
  - 기상청 API 호출 및 데이터 처리.
- **프론트엔드**: Next.js (React 기반)
  - 사용자 인터페이스(UI)를 구성하고, Golang 서버에서 데이터를 받아와 화면에 표시.
- **운영체제**: 라즈베리파이의 경우 Raspbian OS 또는 Ubuntu, 라테판다의 경우 Windows 또는 Linux.

---

## **2. 하드웨어 설정**

### **2.1 라즈베리파이/라테판다 초기 설정**
1. **라즈베리파이를 선택한 경우**:
   - Raspbian OS(또는 Ubuntu)를 microSD 카드에 설치하고 부팅합니다.
   - 초기 설정(네트워크 연결, 업데이트 등)을 완료합니다.
   
   ```bash
   sudo apt update && sudo apt upgrade
   ```

2. **라테판다를 선택한 경우**:
   - Windows 또는 Linux 운영체제를 부팅하고 초기 설정을 완료합니다.
   - 네트워크 연결 및 필수 소프트웨어 설치.

### **2.2 디스플레이 연결**
- HDMI 디스플레이를 보드에 연결하고, 화면 출력을 확인합니다.
- 터치스크린을 사용할 경우 드라이버 설치가 필요할 수 있습니다.

---

## **3. 소프트웨어 개발 환경 구축**

### **3.1 Golang 백엔드 환경 설정**
1. 라즈베리파이나 라테판다에 Golang을 설치합니다.

   ```bash
   wget https://golang.org/dl/go1.19.linux-armv6l.tar.gz  # 최신 버전 다운로드
   sudo tar -C /usr/local -xzf go1.19.linux-armv6l.tar.gz
   export PATH=$PATH:/usr/local/go/bin
   ```

2. Golang 프로젝트 디렉토리를 생성하고, 기상청 API를 호출하는 코드를 작성합니다.

   ```bash
   mkdir weather-backend && cd weather-backend
   go mod init weather-backend
   ```

3. 기상청 API에서 데이터를 가져오는 간단한 HTTP 서버 코드를 작성합니다(앞서 제공한 예시 코드 참고).

4. 서버 테스트:

   ```bash
   go run main.go
   ```

5. 브라우저에서 `http://localhost:8080/weather`로 접속하여 JSON 응답을 확인합니다.

### **3.2 Next.js 프론트엔드 환경 설정**
1. Node.js와 npm을 설치합니다.

   ```bash
   sudo apt install nodejs npm  # Ubuntu/Raspbian 기준
   ```

2. Next.js 프로젝트 생성:

   ```bash
   npx create-next-app weather-display
   cd weather-display
   ```

3. Next.js에서 Golang 백엔드의 `/weather` 엔드포인트로부터 데이터를 가져오는 코드를 작성합니다.
   
   예시로 `pages/index.js` 파일에서 `getServerSideProps`를 사용하여 데이터를 가져옵니다:

   ```javascript
   export async function getServerSideProps() {
     const res = await fetch('http://localhost:8080/weather');
     const data = await res.json();
     return { props: { weatherData: data } };
   }

   export default function Home({ weatherData }) {
     return (
       <div>
         <h1>현관문 앞 날씨</h1>
         <p>온도: {weatherData.temperature}</p>
         <p>습도: {weatherData.humidity}</p>
       </div>
     );
   }
   ```

4. Next.js 개발 서버 실행:

    ```bash
    npm run dev
    ```

5. 브라우저에서 `http://localhost:3000`으로 접속하여 UI가 정상적으로 출력되는지 확인합니다.

---

## **4. 시스템 통합 및 테스트**

### **4.1 백엔드와 프론트엔드 통합**
- Golang 백엔드를 실행한 상태에서 Next.js 프론트엔드를 통해 데이터를 가져오는지 확인합니다.
- 두 서버가 서로 통신할 수 있도록 CORS 설정을 추가하거나, Nginx 같은 리버스 프록시를 설정할 수도 있습니다.

### **4.2 주기적인 데이터 갱신**
- JavaScript의 `setInterval()` 함수를 사용하여 일정 시간마다 날씨 정보를 자동으로 갱신하도록 구현할 수 있습니다.

```javascript
useEffect(() => {
  const interval = setInterval(() => {
    fetchWeather();
  }, 300000); // 5분마다 갱신

  return () => clearInterval(interval);
}, []);
```

---

## **5. 배포 및 최종 설정**

### **5.1 서비스 자동 실행 설정**
라즈베리파이나 라테판다 부팅 시 자동으로 Golang 백엔드와 Next.js 프론트엔드를 실행하도록 설정할 수 있습니다.

1. Golang 서비스 자동 실행:
   
    `/etc/systemd/system/golang-weather.service` 파일 생성:

    ```ini
    [Unit]
    Description=Golang Weather Service

    [Service]
    ExecStart=/usr/local/go/bin/go run /path/to/your/main.go
    Restart=always

    [Install]
    WantedBy=multi-user.target
    ```

2. Next.js 서비스 자동 실행:
   
    `/etc/systemd/system/nextjs-weather.service` 파일 생성:

    ```ini
    [Unit]
    Description=Next.js Weather Service

    [Service]
    ExecStart=/usr/bin/npm run start --prefix /path/to/your/nextjs/project
    Restart=always

    [Install]
    WantedBy=multi-user.target
    ```

3. 서비스 활성화:

    ```bash
    sudo systemctl enable golang-weather.service nextjs-weather.service
    sudo systemctl start golang-weather.service nextjs-weather.service
    ```

---

## 결론

이제 프로젝트를 시작하기 위한 준비가 완료되었습니다! 다음은 요약된 단계입니다:

1. 하드웨어 선택 및 초기 설정 (라즈베리파이/라테판다).
2. Golang 백엔드 개발 (기상청 API 호출).
3. Next.js 프론트엔드 개발 (날씨 정보 표시).
4. 시스템 통합 및 테스트.
5. 자동 실행 설정 및 배포.

이 과정을 따라가면 현관문 앞에서 실시간으로 날씨 정보를 제공하는 IoT 디스플레이 시스템을 성공적으로 구축할 수 있을 것입니다!
