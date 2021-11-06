"""
@author RandySun
@create 2021-08-23-22:36
"""
async def websocket_application(scope, receive, send):

    while True:
        event = await receive()
        if event['type'] == 'websocket.connect':
            await send({
                'type': 'websocket.accept'
            })

        if event['type'] == 'websocket.disconnect':
            break

        if event['type'] == 'websocket.receive':
            print(event['text'])
            import json
            # 收到的内容
            rec=json.loads(event['text'])['message']
            await send({
                'type': 'websocket.send',
                'text': json.dumps({'message':'收到了你的：%s'%rec})
            })