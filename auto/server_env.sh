#Check if exists an .env
FILE=.env
if [[ -f "$FILE" ]]; then
#File exists, removing
    rm $FILE
fi

#Creating an empty .env
touch .env

#Populanting with configs
echo DB_NAME=$DB_NAME >> .env
echo DB_HOST=$DB_HOST >> .env
echo DB_USER=$DB_USER >> .env
echo DB_PORT=$DB_PORT >> .env
echo DB_PASS=$DB_PASS >> .env

#Showing content
cat .env