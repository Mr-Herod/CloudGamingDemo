# kill all services
kill $(sudo ps aux | grep './serverGateway' | tr -s ' '| cut -d ' ' -f 2)
kill $(sudo ps aux | grep './serverNaming' | tr -s ' '| cut -d ' ' -f 2)
kill $(sudo ps aux | grep './serverAccount' | tr -s ' '| cut -d ' ' -f 2)
kill $(sudo ps aux | grep './serverGaming' | tr -s ' '| cut -d ' ' -f 2)
kill $(sudo ps aux | grep './serverRecord' | tr -s ' '| cut -d ' ' -f 2)
# start all services
chmod 777 ~/CloudGamingDemo/resource/programs/ffmpeg_linux

cd ~/CloudGamingDemo/Naming
chmod +x serverNaming
<<<<<<< HEAD
nohup ./serverNaming >> ../log_Naming.out &
=======
./serverNaming >> ../log_Naming.out &
>>>>>>> dc61f2ed7e5c5737e1f1cae0200fb8245060c012
sleep 3

cd ~/CloudGamingDemo/Account
chmod +x serverAccount
<<<<<<< HEAD
nohup ./serverAccount >> ../log_Account.out &
=======
./serverAccount >> ../log_Account.out &
>>>>>>> dc61f2ed7e5c5737e1f1cae0200fb8245060c012
#sleep 2

cd ~/CloudGamingDemo/Gateway
chmod +x serverGateway
<<<<<<< HEAD
nohup ./serverGateway >> ../log_Gateway.out &
=======
./serverGateway >> ../log_Gateway.out &
>>>>>>> dc61f2ed7e5c5737e1f1cae0200fb8245060c012
#sleep 2

cd ~/CloudGamingDemo/Gaming
chmod +x serverGaming
<<<<<<< HEAD
nohup ./serverGaming >> ../log_Gaming.out &
=======
./serverGaming >> ../log_Gaming.out &
>>>>>>> dc61f2ed7e5c5737e1f1cae0200fb8245060c012
#sleep 2

cd ~/CloudGamingDemo/Record
chmod +x serverRecord
<<<<<<< HEAD
nohup ./serverRecord >> ../log_Record.out &
=======
./serverRecord >> ../log_Record.out &
>>>>>>> dc61f2ed7e5c5737e1f1cae0200fb8245060c012


