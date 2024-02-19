#pragma once
#include "constant.h"
#include "user.h"
#include <vector>
class ChatRoom {
public:
	static inline ChatRoom* GetInstance();
	static inline void AddUser(const User&user);
	static inline void DelUser(const User&user);
private:
	ChatRoom(){}
	ChatRoom(const ChatRoom& rval) = delete;
	ChatRoom operator=(ChatRoom val) = delete;
private:
	static int max_user_num;
	static int cur_user_num;
	static std::vector<User> user_arr;
};