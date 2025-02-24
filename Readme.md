# Go Fiber Library API
This is a library API built with the Go programming language and the [Fiber](https://github.com/gofiber/fiber) framework. It demonstrates how to create a RESTful API to manage book in library 

## Technologie(s)
- Go (1.23.5)
- Go Modules
- MySQL (or other databases of your choice)

### Description:
1. **Create a book (`POST /book`)**: ใช้สำหรับสร้างข้อมูลหนังสือ
2. **Get All Books (`GET /books`)**:                  
>     ใช้เพื่อดึงข้อมูลหนังสือทั้งหมด
>     ใช้ข้อมูลหนังสือทั้งหมด โดยสามารถค้นหาจากชื่อหนังสือ, ผู้แต่ง, หมวดหมู่
>     ใช้เพื่อดึงข้อมูลจำนวนการยืมหนังสือของเเต่ละเล่ม เเล้วสามารถเรียงลำดับของข้อมูล
3. **Get Book by id (`GET /book/{id}`)**: ใช้เพื่อดึงข้อมูลของหนังสือตาม id ที่ระบุ
4. **Update Book (`PUT /book`)**: ใช้เพื่ออัปเดต/เเก้ไขข้อมูลหนังสือที่มีอยู่แล้ว
5. **Update Book status (`PUT /book/status`)**: 
>     ใช้เพื่อยืม/คืนหนังสือข้อมูลหนังสือที่มีอยู่แล้ว
>     ลงข้อมูล log ในการทำ transaction ยืม/คืนหนังสือ
6. **Delete Book (`DELETE /book/{id}`)**: ใช้สำหรับลบหนังสือตาม id ที่ระบุ

### Information:
- ใช้สำหรับทดสอบการเรียก API Library
- ใช้เครื่องมืออย่าง Postman เพื่อทดสอบ API