"""
ASGI config for chatsite project.

It exposes the ASGI callable as a module-level variable named ``application``.

For more information on this file, see
https://docs.djangoproject.com/en/3.0/howto/deployment/asgi/
"""

import os
import django
from django.core.asgi import get_asgi_application
from chat.websocket import websocket_application

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'chatsite.settings')
django.setup()
application_asgi = get_asgi_application()
async def application(scope, receive, send):
    print(scope)
    if scope['type'] == 'http':
        print('http')
        await application_asgi(scope, receive, send)
    elif scope['type'] == 'websocket':
        print('websocket')
        await websocket_application(scope, receive, send)
    else:
        raise NotImplementedError("Unknown scope type %s"%scope['type'])