package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	//1 找int切片中最大的元素(不准用排序)
	var (
		j int
		x int
	)
	nums := []int{3, 11, 5, 7, 9}
	j = nums[0] //赋予j初始值
	for _, v := range nums {
		//比较值
		if v > j {
			j = v
		}
		nums = nums[1:] //利用队列特性，比较一次后减去一个元素
	}
	fmt.Printf("最大的元素为%d\n", j)

	//2 找int切片中第二个最大元素(不准用排序)
	nums = []int{100, 3, 30, 21, 9, 10, 20}
	j = nums[0] //赋予j初始值
	x = nums[1] //赋予x初始值
	for i := 2; i < len(nums); i++ { //跳过初始赋值索引进行比较
		if nums[i] > x && nums[i] < j { //如果元素比x大且元素比j小
			x = nums[i] //第二大
		} else {
			j = nums[i] //最大值
		}
	}
	fmt.Printf("第二大的元素为%d\n", x)

	//3 获取映射中所有key组成的切片/获取所有value组成的切片
	numb := map[string]int{"zjj": 8, "jk": 9, "jack": 10}
	keysilce := []string{}
	valuesilce := []int{}
	for k, v := range numb {
		keysilce = append(keysilce, k)     //获取key的切片
		valuesilce = append(valuesilce, v) //获取value的切片
	}
	fmt.Println(keysilce, valuesilce)

	//4 我有一个梦想
	dream := `
	I am happy to join with you today in what will go down in history as the greatest demonstration for freedom in the history of our nation.

		Five score years ago, a great American, in whose symbolic shadow we stand today, signed the Emancipation Proclamation. This momentous decree came as a great beacon light of hope to millions of Negro slaves who had been seared in the flames of withering injustice. It came as a joyous daybreak to end the long night of bad captivity.

		But one hundred years later, the Negro still is not free. One hundred years later, the life of the Negro is still sadly crippled by the manacles of segregation and the chains of discrimination. One hundred years later, the Negro lives on a lonely island of poverty in the midst of a vast ocean of material prosperity. One hundred years later, the Negro is still languished in the corners of American society and finds himself an exile in his own land. And so we've come here today to dramatize a shameful condition.

		In a sense we've come to our nation's capital to cash a check. When the architects of our republic wrote the magnificent words of the Constitution and the Declaration of Independence, they were signing a promissory note to which every American was to fall heir. This note was a promise that all men, yes, black men as well as white men, would be guaranteed the "unalienable Rights" of "Life, Liberty and the pursuit of Happiness." It is obvious today that America has defaulted on this promissory note, insofar as her citizens of color are concerned. Instead of honoring this sacred obligation, America has given the Negro people a bad check, a check which has come back marked "insufficient funds."

	But we refuse to believe that the bank of justice is bankrupt. We refuse to believe that there are insufficient funds in the great vaults of opportunity of this nation. And so, we've come to cash this check, a check that will give us upon demand the riches of freedom and the security of justice.

		We have also come to this hallowed spot to remind America of the fierce urgency of Now. This is no time to engage in the luxury of cooling off or to take the tranquilizing drug of gradualism. Now is the time to make real the promises of democracy. Now is the time to rise from the dark and desolate valley of segregation to the sunlit path of racial justice. Now is the time to lift our nation from the quicksands of racial injustice to the solid rock of brotherhood. Now is the time to make justice a reality for all of God's children.

		It would be fatal for the nation to overlook the urgency of the moment. This sweltering summer of the Negro's legitimate discontent will not pass until there is an invigorating autumn of freedom and equality. Nineteen sixty-three is not an end, but a beginning. And those who hope that the Negro needed to blow off steam and will now be content will have a rude awakening if the nation returns to business as usual. And there will be neither rest nor tranquility in America until the Negro is granted his citizenship rights. The whirlwinds of revolt will continue to shake the foundations of our nation until the bright day of justice emerges.

		But there is something that I must say to my people, who stand on the warm threshold which leads into the palace of justice: In the process of gaining our rightful place, we must not be guilty of wrongful deeds. Let us not seek to satisfy our thirst for freedom by drinking from the cup of bitterness and hatred. We must forever conduct our struggle on the high plane of dignity and discipline. We must not allow our creative protest to degenerate into physical violence. Again and again, we must rise to the majestic heights of meeting physical force with soul force.

		The marvelous new militancy which has engulfed the Negro community must not lead us to a distrust of all white people, for many of our white brothers, as evidenced by their presence here today, have come to realize that their destiny is tied up with our destiny. And they have come to realize that their freedom is inextricably bound to our freedom.

		We cannot walk alone.

		And as we walk, we must make the pledge that we shall always march ahead.

		We cannot turn back.

		There are those who are asking the devotees of civil rights, "When will you be satisfied?" We can never be satisfied as long as the Negro is the victim of the unspeakable horrors of police brutality. We can never be satisfied as long as our bodies, heavy with the fatigue of travel, cannot gain lodging in the motels of the highways and the hotels of the cities. We cannot be satisfied as long as the Negro's basic mobility is from a smaller ghetto to a larger one. We can never be satisfied as long as our children are stripped of their selfhood and robbed of their dignity by signs stating "for whites only." We cannot be satisfied as long as a Negro in Mississippi cannot vote and a Negro in New York believes he has nothing for which to vote. No, no, we are not satisfied, and we will not be satisfied until "justice rolls down like waters, and righteousness like a mighty stream."

	I am not unmindful that some of you have come here out of great trials and tribulations. Some of you have come fresh from narrow jail cells. And some of you have come from areas where your quest -- quest for freedom left you battered by the storms of persecution and staggered by the winds of police brutality. You have been the veterans of creative suffering. Continue to work with the faith that unearned suffering is redemptive. Go back to Mississippi, go back to Alabama, go back to South Carolina, go back to Georgia, go back to Louisiana, go back to the slums and ghettos of our northern cities, knowing that somehow this situation can and will be changed.

		Let us not wallow in the valley of despair, I say to you today, my friends.

		And so even though we face the difficulties of today and tomorrow, I still have a dream. It is a dream deeply rooted in the American dream.

		I have a dream that one day this nation will rise up and live out the true meaning of its creed: "We hold these truths to be self-evident, that all men are created equal."

	I have a dream that one day on the red hills of Georgia, the sons of former slaves and the sons of former slave owners will be able to sit down together at the table of brotherhood.

		I have a dream that one day even the state of Mississippi, a state sweltering with the heat of injustice, sweltering with the heat of oppression, will be transformed into an oasis of freedom and justice.

		I have a dream that my four little children will one day live in a nation where they will not be judged by the color of their skin but by the content of their character.

		I have a dream today!

		I have a dream that one day, down in Alabama, with its vicious racists, with its governor having his lips dripping with the words of "interposition" and "nullification" -- one day right there in Alabama little black boys and black girls will be able to join hands with little white boys and white girls as sisters and brothers.

		I have a dream today!

		I have a dream that one day every valley shall be exalted, and every hill and mountain shall be made low, the rough places will be made plain, and the crooked places will be made straight; "and the glory of the Lord shall be revealed and all flesh shall see it together."

	This is our hope, and this is the faith that I go back to the South with.

		With this faith, we will be able to hew out of the mountain of despair a stone of hope. With this faith, we will be able to transform the jangling discords of our nation into a beautiful symphony of brotherhood. With this faith, we will be able to work together, to pray together, to struggle together, to go to jail together, to stand up for freedom together, knowing that we will be free one day.

		And this will be the day -- this will be the day when all of God's children will be able to sing with new meaning:

	My country 'tis of thee, sweet land of liberty, of thee I sing.

		Land where my fathers died, land of the Pilgrim's pride,

		From every mountainside, let freedom ring!

		And if America is to be a great nation, this must become true.

		And so let freedom ring from the prodigious hilltops of New Hampshire.

		Let freedom ring from the mighty mountains of New York.

		Let freedom ring from the heightening Alleghenies of

	Pennsylvania.

		Let freedom ring from the snow-capped Rockies of Colorado.

		Let freedom ring from the curvaceous slopes of California.

		But not only that:

	Let freedom ring from Stone Mountain of Georgia.

		Let freedom ring from Lookout Mountain of Tennessee.

		Let freedom ring from every hill and molehill of Mississippi.

		From every mountainside, let freedom ring.

		And when this happens, when we allow freedom ring, when we let it ring from every village and every hamlet, from every state and every city, we will be able to speed up that day when all of God's children, black men and white men, Jews and Gentiles, Protestants and Catholics, will be able to join hands and sing in the words of the old Negro spiritual:

	Free at last! Free at last!

		Thank God Almighty, we are free at last!
		`

	//统计大小写英文字母次数
	//使用count作为key
	//value rune
	//for range
	//ascii ..A-Z ..a-z...
	//按照大小写字母顺序输出统计次数
	nums_dream := map[int][]rune{} //定义空映射
	nums_sort := []int{}           //定义切片对key值进行排序
	for _, i := range dream { //遍历变量内容
		if (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
			//fmt.Println(nums_dream[1][0])
			//fmt.Println(strings.Count(dream,string(i)),i)
			if strings.Contains(string(nums_dream[strings.Count(dream, string(i))]), string(i)) == true { //判断切片对应的key内容是否已包含追加的值，如包含则忽略（切片去重)
				continue
			} else {
				nums_dream[strings.Count(dream, string(i))] = append(nums_dream[strings.Count(dream, string(i))], i) //根据相同的key追加value到切片
			}
		}
	}
	for k, _ := range nums_dream {
		nums_sort = append(nums_sort, int(k)) //将映射中的key存放到切片中
		//fmt.Printf("%d:%c\n",k,v)  //通过格式化将码点转换成字符char
	}

	//fmt.Println(nums_sort, nums_dream)
	sort.Ints(nums_sort) //对key值进行排序
	for _, v := range nums_sort {
		fmt.Printf("%d:%c\n", v, nums_dream[v]) //使用排好序的key进行输出并且获取映射中的value
	}

	//冒泡排序
	heght := []int{10, 6, 7, 9, 5}
	//遍历
	for n := 0; n < len(heght)-1; n++ { //重复内部元素比较次数
		for i := 0; i < len(heght)-1; i++ { //内部元素比较，逐个元素进行比较，将最大的数移动到最后
			if heght[i] > heght[i+1] {
				heght[i], heght[i+1] = heght[i+1], heght[i]
			}
		}
	}
	fmt.Println(heght)

	//插入排序
	/*
		从第一个元素开始，该元素可以认为已经被排序
		取出下一个元素，在已经排序的元素序列中从后向前扫描
		如果该元素（已排序）大于新元素，将该元素移到下一位置
		重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
		将新元素插入到该位置后
		重复步骤2~5
	*/
	nums = []int{16, 5, 3, 56, 7} //5<16;n=0;j=5;n>=0 && 16>5;0:5,1:16;{5,16,3,56,7};[16 5 3 56 7][5 16 3 56 7][3 5 16 56 7][3 5 7 16 56]
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			n := i - 1   //记录前一个元素位置
			j := nums[i] //记录当前值
			for n >= 0 && nums[n] > j { //使用n作为条件，判断前一个元素值是否大于当前值，如是从当前位置进行交换(一直往前进行比较交换直到循环结束)
				nums[n+1], nums[n] = nums[n], nums[n+1] //交换位置
				n--
			}
		}
	}
	fmt.Println(nums)

	//猜数字（二分查找)
	guess := []int{}
	//var k int
	for j := 1; j <= 100; j++ { //添加一个1-100的切片
		guess = append(guess, j)
	}
	//fmt.Printf("%#v\n",guess)
	fmt.Println("开始生成100以内的随机数...")
	rand.Seed(time.Now().Unix())
	guess_num := rand.Intn(100) //随机生成100以内的数
	i := sort.Search(len(guess), func(i int) bool {return guess[i] >= guess_num}) //通过查找将index赋值给变量i
	fmt.Printf("查找结果:在切片中查询出来的位置是%d,切片值:%d\t随机生成的数字为:%d",i,guess[i],guess_num)

}
