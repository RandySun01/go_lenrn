"""
@author RandySun
@create 2021-08-22-16:19
"""

from django.urls import path
from . import consumers
from channels.auth import AuthMiddlewareStack
from channels.routing import ProtocolTypeRouter, URLRouter
from django.urls import re_path

application = ProtocolTypeRouter(
    {
        "websocket": AuthMiddlewareStack(
            URLRouter([
                re_path(r'ws/chat/(?P<room_name>\w+)/$', consumers.ChatConsumer),

            ])
        )
    }
)
