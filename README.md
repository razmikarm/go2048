# go2048
2048 game in CLI written on Go


##### Make sure you have permission to run
>chmod +x go2048

### How to play

> Use arrow keys to move points

#### Simple run
> ./go2048

![image](https://user-images.githubusercontent.com/54362304/166344564-fab67aab-42cb-454b-90fa-2620f25e8972.png)

![image](https://user-images.githubusercontent.com/54362304/166695330-4f6160bc-d9e9-4621-acc6-d4e3a38d20c3.png)


#### Filling border with alternative symbol
> ./go2048 "*"

![image](https://user-images.githubusercontent.com/54362304/166344533-1ff87920-83c2-4b07-8aa5-fcb7e79df6e9.png)

You can fill borders with any character you want. Just run `./go2048 <char>`. Here we set `*` in double quotes because in shell it is a special character, like `|`, `;`, `&`, `<`, `>` and etc. For the same reasons in place of `<char>` use `"\""` for `"` and `"\\"` for `\`.
