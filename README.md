## 콘셉가정
- 노티 수단은 계속 추가될 수 있다
    - 특정 노티 수단과 관련된 테이블이 있다면 테이블이 계속 추가되므로
      NotiMethod 테이블의 `type` 과 `identifier` 로 관리
    - 현재 지원하는 노티 수단 type: `email`, `slack`, `sms`

- 1명의 유저가 1개 이상의 동일 노티수단을 가질 수 있다.
- (1개 이상의 email 과 1개 이상의 slack 채널로 노티를 받을 수 있다.)
    - user service 에서 가지고 있는 email 외의 다른 email 을 등록, verify 후
      노티 받을 수 있음. 여러 slack 채널id 로 노티 받을 수 있음.
    - 단 유저의 기본 email 은 저장하지 않고,
      `{ type: 'email', identifier: 'default' }` 로 계정 기본 email 을 사용.

- fail 알림은 `endpoint_id` 


## 다른 서비스와 협의 해야할 사항
- user service: user 의 UTC offset 을 프론트엔드에 제공해야함

- endpoint service: 
    - fail event 발생하여 api 호출시 파라미터로 `endpoint_id` 필요
    - `endpoint_id` 필요한 이유:
        - userA 가 a.com/ping 을 등록
        - userB 가 a.com/ping 을 등록
        - a.com 의 ping api 주소가 `/ping` 에서 `/helthcheck` 으로 변경됨
        - userA 가 a.com/ping -> a.com/helthcheck 으로 변경
        - userB 는??

        - a.com/ping 이라는 url 을 endpoint 객체로 한번 추상화
        ```
        {
          id: '50f8d91',
          name: 'Our a.com service!',
          description: 'a.com monitoring',
          url: 'http://a.com',
          healthCheckUrl: 'http://a.com/ping',
        }
        ```
        - userA 가 `50f8d91` 의 `healthCheckUrl` 을 변경시 userB 도 변경된
          주소로 잘 받게됨
        - ( 이때 endpoint 정보 change noti 도 필요함!)


- email sender, sms sender, slack sender services:
    - email address, phone number, slack channel Id 등을 가지고 있지 않으면 좋을듯.
        - emailSender 서비스를 호출할때 보내는 파라미터 예:
        ```
        { endpointId, email, type }
        ```

    - healthCheck 체크 실패 이외의 type 추가(각 타입에 대한 템플릿 필요):
        - `endpoint-fail`: endpoint helth-check 실패시.
        - `endpoint-chagne`: endpoint 정보 변경시.
        - `noti-method-verify`: 추가 노티 수단 등록시 verify 용도. 각 파라미터는
            - email 일땐 `{ userid, email, comfirmLink }`
            - slack 일땐 `{ ??? }`
            - sms 일땐 `{ userid, phoneNumber, code }`

## API

### Case1: 유저가 엔드포인트를 등록했을때,
POST /v1/endpointSetting/:userId/:endpointId
insert record into UserEndpointSetting

### Case2: 유저가 엔드포인트를 삭제했을때,
DELETE /v1/endpointSetting/:userId/:endpointId
delete record from UserEndpointSetting

### Case3: 유저가 특정 엔드포인트의 알람 설정을 변경할때,
PUT /v1/endpointSetting/:userId/:endpointId
request parameter: { json }
update record in UserEndpointSetting

### Case4: 유저의 global noti 설정 조회
GET /v1/globalSetting/:userId

### Case5: 유저의 모든 specific-endpoint noti 설정 조회
GET /v1/endpointSetting/:userId

### Case5: 유저의 specific-endpoint noti 설정 조회
GET /v1/endpointSetting/:userId/:endpointId

### Case6: 특정 엔드포인트 실패 알림을 받았을때,
POST /v1/fail/:endpointId

### Case7: 특정 엔드포인트 정보 변경 알림,
PUT /v1/endpointSetting/:endpointId


## json
https://mholt.github.io/json-to-go/
