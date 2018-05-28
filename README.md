# gb2260

## latest revision
- 2018.02

## source list
- 201802	http://www.mca.gov.cn/article/sj/xzqh/2018/201804-12/201804121005.html	
- 201712	http://www.mca.gov.cn/article/sj/tjbz/a/2018/201803131439.html	
- 201612	http://www.mca.gov.cn/article/sj/tjbz/a/2016/201612/201705311652.html	
- 201512	http://www.mca.gov.cn/article/sj/tjbz/a/2015/201706011127.html	
- 201412	http://files2.mca.gov.cn/cws/201502/20150225163817214.html	
- 201312	http://files2.mca.gov.cn/cws/201404/20140404125552372.htm	
- 201212	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201707271556.html	
- 201112	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201707271552.html	
- 201012	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220946.html	
- 200912	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220943.html	
- 200812	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220941.html	
- 200712	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220939.html	
- 200612	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220936.html	
- 200512	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220935.html	
- 200412	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220930.html	
- 200312	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220928.html	
- 200212	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220927.html	
- 200112	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220925.html	
- 200012	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220923.html	
- 199912	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220921.html	
- 199812	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220918.html	
- 199712	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220916.html	
- 199612	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220914.html	
- 199512	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220913.html	
- 199412	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220911.html	
- 199312	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041023.html	
- 199212	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220910.html	
- 199112	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041020.html	
- 199012	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041018.html	
- 198912	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041017.html	
- 198812	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220903.html	
- 198712	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220902.html	
- 198612	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220859.html	
- 198512	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220858.html	
- 198412	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220856.html	
- 198312	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708160821.html	
- 198212	http://www.mca.gov.cn/article/sj/tjbz/a/1980-2000/201707141125.html	
- 198112	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708041004.html	
- 198012	http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708040959.html

## the src file dir
- data/

## the go file dir
- gbdata/

## customize
you can modify the __generate.go__ to create your own format

# tsv file generate script
https://github.com/cn/GB2260
you can midify the sources.tsv and use `make` to generate the src data

# usage
if you use golang as your language, you can use https://github.com/vgmdj/utils as your parsing package directly

```
package main

import(
    "fmt"
    "github.com/vgmdj/utils/area"
    _ "github.com/vgmdj/gb2260/gbdata"  //导入数据包
)

func main(){
   gb2260 := area.NewArea(area.GB2260)
   bj := gb2260.Get("110101")
   fmt.Println(bj.Province)    //北京市
   fmt.Println(bj.County)      //东城区
   fmt.Println(bj.FullName())  //北京市东城区
}


```
