import cpipe
cpipe.add_command("hset","words","word1","1")
cpipe.add_command("hset","words","word2","1")
cpipe.add_command("hset","words","word3","1")
cpipe.add_command("hset","words","word4","1")
cpipe.add_command("hset","words","word5","1")
cpipe.add_command("hset","words","word6","1")
print(cpipe.execute(1))
