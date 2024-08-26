# ใช้ภาพ MongoDB ล่าสุดจาก Docker Hub
FROM mongo:latest

# กำหนดพอร์ตที่ MongoDB จะทำงาน
EXPOSE 27017

# กำหนดโฟลเดอร์ที่จะใช้เก็บข้อมูลของ MongoDB
VOLUME /data/db

# คำสั่งเริ่มต้นในการรัน MongoDB
CMD ["mongod"]
