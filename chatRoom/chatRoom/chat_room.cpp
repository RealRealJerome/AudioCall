#include"chat_room.h"
ChatRoom* ChatRoom::GetInstance() {
	static ChatRoom chat_room;
	return &chat_room;
}
void ChatRoom::AddUser(const User& user) {
	user_arr.push_back(user);
}
void ChatRoom::DelUser(const User& user) {
	user_arr.erase(std::find(user_arr.begin(), user_arr.end(), user));
}