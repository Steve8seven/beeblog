{{define "navbar"}}
<a class="navbar-brand" href="/">Stevano的博客</a>
<div>
    <ul class="nav nav-pills" style="margin-top: 5px;">
       <li {{if .isHome}}class="active"{{end}}><a href="/">主页</a></li>
       <li {{if .isCategory}}class="active"{{end}}><a href="/category">分类</a></li>
       <li {{if .isTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
       {{if .IsLogin}}
       <li class="pull-right"><a href="/login?isExit=true">退出</a></li>
       {{else}}
       <li class="pull-right"><a href="/login">管理员登录</a></li>
       {{end}}
    </ul>
</div>
{{end}}