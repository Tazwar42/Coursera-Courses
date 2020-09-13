from django.urls import path
from django.views.generic import TemplateView
from . import views

app_name = 'hello'

urlpatterns = [
    path('', views.sessfun),
    path('cookie', views.cookie),
    path('sessfun', views.sessfun),
]