#pragma once
#include<string>
#include"global.h"
class User {
public:
	User(std::string user_name, std::string ip, std::string enter_time, std::string leave_time) :user_name(user_name), ip(ip), enter_time(enter_time), leave_time(leave_time) {}
	bool operator==(const User& val) const{ return this->user_name == val.user_name; }
	void BroadCast();
private:
	std::string user_name; // �û��� -- Ψһ��ʶ
	std::string ip; // �û���ip
	std::string enter_time; // ����ʱ��
	std::string leave_time; // �뿪ʱ��
};