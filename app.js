var request = require("request");

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
  "&" + encodeURIComponent("base_time") + "=" + encodeURIComponent("0500");
queryParams += "&" + encodeURIComponent("nx") + "=" + encodeURIComponent("35");
queryParams += "&" + encodeURIComponent("ny") + "=" + encodeURIComponent("128");
queryParams += "&" + encodeURIComponent("authKey") + "=wUFFWyipQ6CBRVsoqROgzg";

let contents = "";

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
        console.log(B.response.body.items.item[i].fcstValue);
      }
    }
  }
);
