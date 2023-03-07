package main

func main() {
router := gin.New()
router.Get("/", func(c *gin.Context){
c.String(200, "Hello World!")
})
router.Run(":8080")
}