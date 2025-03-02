import React from 'react';
import { Layout, Menu } from 'antd';
import {  useLocation } from 'react-router-dom';
import { SidebarMenuItems } from '@/config';
const { Sider } = Layout;

const SideMenu = () => {
  const location = useLocation();

  return (
    <Sider style={{ zIndex: 1000 }}>
      <div className="demo-logo-vertical">
         ADMIN PANEL
      </div>
      <Menu
        theme="dark"
        mode="inline"
        defaultSelectedKeys={['/dashboard']}
        selectedKeys={[location.pathname]}
        items={SidebarMenuItems}
      />
    </Sider>
  );
};

export default SideMenu; 