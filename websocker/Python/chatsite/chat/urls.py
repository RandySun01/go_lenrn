"""
@author RandySun
@create 2021-08-22-15:48
"""
# chat/urls.py
from django.urls import path

from . import views

urlpatterns = [
    path('', views.index, name='index'),
    path('<str:room_name>/', views.room, name='room'),
    path('asgi/', views.index1),
]