#pragma once
#include<string>
#include"global.h"
class User {
public:
	User(std::string user_name, std::string ip, std::string enter_time, std::string leave_time) :user_name(user_name), ip(ip), enter_time(enter_time), leave_time(leave_time) {}
	bool operator==(const User& val) const{ return this->user_name == val.user_name; }
	void BroadCast();
private:
	std::string user_name; // 用户名 -- 唯一标识
	std::string ip; // 用户的ip
	std::string enter_time; // 加入时间
	std::string leave_time; // 离开时间
};