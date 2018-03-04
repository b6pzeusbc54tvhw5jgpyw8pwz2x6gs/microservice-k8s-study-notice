
## 콘셉가정
- 노티 수단은 계속 추가될 수 있다
    - 특정 노티 수단과 관련된 테이블이 있다면 테이블이 계속 추가되므로
      NotiMethod 테이블의 `type` 과 `identifier` 로 관리
    - 현재 지원하는 노티 수단 type: `email`, `slack`, `sms`

- 1명의 유저가 1개 이상의 email 과 1개 이상의 slack 채널로 노티를 받을 수 있다.
    - user service 에서 가지고 있는 email 외의 email 이나, 여러 slack 채널id 를
      가지고 있어야함.
    - { type: 'email', identifier: 'default' } 인 경우는 user service 의 email
      을 사용

- fail 알림은 `endpoint_id` 

## 기본전제 process
1. endpoint check 실패시, endpoint id 전달
1. UserEndpointSetting 테이블에서 User 들을 모두 찾음
1. U


## Case1: 새로운 유저 생성 노티를 받았을때,
아무것도 안함

## Case2: 유저가 엔드포인트를 등록했을때,
insert record into UserEndpointSetting

## Case3: 유저가 엔드포인트를 삭제했을때,
delete record from UserEndpointSetting

## Case4: 유저가 특정 엔드포인트의 알람 설정을 변경할때,
update record in UserEndpointSetting

## Case5: 특정 엔드포인트 실패 알림을 받았을때,
1. request parameter 로 넘어온 `endpoint_id` 로 UserEndpointSetting 테이블을 조회화여 `user_id` 를 모두 찾는다.
2. 

## Case4: a user change own global setting
update(or insert) record in UserSetting

## Case5: a user change own notiType setting
update(or insert) record in UserNotiTypeSetting

## Case6: a user change own endpoint setting
update(or insert) record in UserNotiTypeSetting

## 다른 서비스에 요청 할 것
- userService: user 의 UTC offset 을 프론트엔드에 제공해야함

- endpoint service: 

    - `endpoint_id` 필요. why?
        - user A 가 a.com/ping 을 등록
        - user B 가 a.com/ping 을 등록
        - a.com 의 ping api 주소가 /v1/ping 으로 변경됨
        - user A 가 a.com/ping -> a.com/v1/ping 으로 변경
        - user B 는??

        - a.com/ping 이라는 url 을 endpoint 객체로 한번 추상화
        - `{ id: '', url: 'http://a.com', healthCheckUrl: 'http://a.com/ping'
    
    url 로 관리한다면 ping api 주소가 변경되었을때



    - fail event 발생하여 api 호출시 파라미터로 `endpoint_id` 필요

- email sender, sms sender, slack sender services: email address, phone number, slack channel Id 등을 가지고 있지 않으면 좋을듯. email 전송 필요시 제공해주는 api 로
`{ endpointId, email }` 만 보낼 것임


## 유저가 email 변경을 하면?
,

