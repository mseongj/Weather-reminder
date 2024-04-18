import { createRequire } from "module";
const require = createRequire(import.meta.url);
// require("request").config();

var request = require("request");
import config from "./apiKey.js";
const { API_KEY } = config;

let today = new Date();

let year = today.getFullYear();
let month = ("0" + (1 + today.getMonth())).slice(-2);
let date = ("0" + today.getDate()).slice(-2);

let DAY = `${year}${month}${date}`;

// 개황	대구	143
// 육상	대구	11H10701
// 문 앞 날씨를 알려주는 기기에 날씨 정보를 확인 하기 위한 api신청입니다. (상업적 사용X, 사용되는 코드는 GitHub에 올라갈 예정 입니다.)
// https://apihub.kma.go.kr/api/typ02/openApi/VilageFcstInfoService_2.0/getVilageFcst?pageNo=1
// &numOfRows=12
// &dataType=JSON
// &base_date=20240111
// &base_time=0500
// &nx=35
// &ny=128
// &authKey=wUFFWyipQ6CBRVsoqROgzg

let url =
  "https://apihub.kma.go.kr/api/typ02/openApi/VilageFcstInfoService_2.0/getVilageFcst";
let queryParams =
  "?" + encodeURIComponent("pageNo") + "=" + encodeURIComponent("1");
queryParams +=
  "&" + encodeURIComponent("numOfRows") + "=" + encodeURIComponent("240");
queryParams +=
  "&" + encodeURIComponent("dataType") + "=" + encodeURIComponent("JSON");
queryParams +=
  "&" + encodeURIComponent("base_date") + "=" + encodeURIComponent(DAY);
queryParams +=
  "&" + encodeURIComponent("base_time") + "=" + encodeURIComponent("1700");
queryParams += "&" + encodeURIComponent("nx") + "=" + encodeURIComponent("36");
queryParams += "&" + encodeURIComponent("ny") + "=" + encodeURIComponent("127");
queryParams += "&" + encodeURIComponent("authKey") + "=" + `${API_KEY}`;

let tempture = [];
let uuu = [];
let vvv = [];
let vec = [];
let wsd = [];
let sky = [];
let pty = [];
let pop = [];
let wav = [];
let pcp = [];
let reh = [];
let sno = [];

let i = 0;
request(
  {
    url: url + queryParams,
    method: "GET",
  },
  (error, response, contents) => {
    console.log(
      "error status: ",
      error === null ? "error is not exeist" : error
    );
    console.log("Status", response.statusCode);
    let B = JSON.parse(contents);
    // console.log("Headers", response.headers);
    // console.log("Reponse received", B);
    // console.log(B.response.body.items.item[0].category);
    for (i = 0; i < 240; i++) {
      if (B.response.body.items.item[i].category == "TMP") {
        tempture.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "UUU") {
        uuu.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "VVV") {
        vvv.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "VEC") {
        vec.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "WSD") {
        wsd.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "SKY") {
        sky.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "PTY") {
        pty.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "POP") {
        pop.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "WAV") {
        wav.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "PCP") {
        pcp.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "REH") {
        reh.push(B.response.body.items.item[i].fcstValue);
      }
      if (B.response.body.items.item[i].category == "SNO") {
        sno.push(B.response.body.items.item[i].fcstValue);
      }
    }
    console.log("1시간 기온: ", tempture);
    console.log("풍속 (동서): ", uuu);
    console.log("풍속 (남북): ", vvv);
    console.log("풍향: ", vec);
    console.log("풍속: ", wsd);
    console.log("하늘상태: ", sky);
    console.log("강수형태: ", pty);
    console.log("강수확률: ", pop);
    console.log("파고: ", wav);
    console.log("1시간 강수량: ", pcp);
    console.log("습도: ", reh);
    console.log("1시간 신적설: ", sno);
  }
);

// 하늘상태 (SKY)코드: 맑음(1), 구름많음(3), 흐림(4)
// 강수형태 (PTY)코드: 없음(0) 비(1), 비/눈(2), 눈(3), 빗방울(5), 빗방울/눈날림(6), 눈날림(7)

// 01-12
