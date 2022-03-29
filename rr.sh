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
./serverNaming >> ../log_Naming.out &
sleep 3

cd ~/CloudGamingDemo/Account
chmod +x serverAccount
./serverAccount >> ../log_Account.out &
#sleep 2

cd ~/CloudGamingDemo/Gateway
chmod +x serverGateway
./serverGateway >> ../log_Gateway.out &
#sleep 2

cd ~/CloudGamingDemo/Gaming
chmod +x serverGaming
./serverGaming >> ../log_Gaming.out &
#sleep 2

cd ~/CloudGamingDemo/Record
chmod +x serverRecord
./serverRecord >> ../log_Record.out &


