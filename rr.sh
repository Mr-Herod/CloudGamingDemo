# kill all services
kill $(sudo ps aux | grep './serverGateway' | tr -s ' '| cut -d ' ' -f 2)
kill $(sudo ps aux | grep './serverNaming' | tr -s ' '| cut -d ' ' -f 2)
kill $(sudo ps aux | grep './serverAccount' | tr -s ' '| cut -d ' ' -f 2)
kill $(sudo ps aux | grep './serverGaming' | tr -s ' '| cut -d ' ' -f 2)
kill $(sudo ps aux | grep './serverRecord' | tr -s ' '| cut -d ' ' -f 2)
# start all services
cd ~/CloudGamingDemo/Naming
chmod +x serverNaming
./serverNaming >> log.out &
sleep 2

cd ~/CloudGamingDemo/Account
chmod +x serverAccount
./serverAccount >> log.out &
sleep 2

cd ~/CloudGamingDemo/Gateway
chmod +x serverGateway
./serverGateway >> log.out &
sleep 2

cd ~/CloudGamingDemo/Gaming
chmod +x serverGaming
./serverGaming >> log.out &
sleep 2

cd ~/CloudGamingDemo/Record
chmod +x serverRecord
./serverRecord >> log.out &


