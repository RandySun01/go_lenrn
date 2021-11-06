from django.shortcuts import render

# Create your views here.
from django.shortcuts import render,HttpResponse
import queue

# Create your views here.



# 全局定义一个字典用来存储客户端浏览器与队列关系
q_dict = {}



def home(request, *args, **kwargs):
    # 获取用户唯一标识
    name = request.GET.get('name')
    # 给每个客户端浏览器创建独有的队列
    q_dict[name] = queue.Queue()
    return render(request,'index.html',locals())


def send_msg(request, *args, **kwargs):
    if request.method == 'POST':
        con = request.POST.get('content')
        name = request.POST.get('name')
        # 将该消息往所有的群聊的客户端队列中添加
        content = {'name': name, 'content': con, }
        for q in q_dict.values():
            q.put(content)
        return HttpResponse('OK')

import json
from django.http import JsonResponse
def get_msg(request, *args, **kwargs):
    name = request.GET.get("name")
    # 去对应的队列中获取数据
    q = q_dict.get(name)
    back_dic = {'status':True,'msg':''}
    try:
        data = q.get(timeout=10)  # 等10s
        back_dic['msg'] = data
    except queue.Empty as e:
        back_dic['status'] = False
    # return HttpResponse(json.dumps(back_dic))
    return JsonResponse(back_dic)
# 大公司一般情况下都会使用上面长轮询的方式，因为兼容性好