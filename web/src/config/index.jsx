import React from 'react';
import { Link } from 'react-router-dom';

import {
    ProfileOutlined,
    SettingOutlined,
    LogoutOutlined,
    DashboardOutlined,
    UserOutlined,
    TeamOutlined,
  } from '@ant-design/icons';

export const HeaderMenuItems = [
    {
      key: '1',
      label: 'Profile',
      icon: <ProfileOutlined />,
    },
    {
      key: '2',
      label: 'Settings',
      icon: <SettingOutlined />,
    },
    {
      type: 'divider',
    },
    {
      key: '3',
      label: 'Logout',
      icon: <LogoutOutlined />,
      danger: true,
    }
  ];

export const SidebarMenuItems = [
    {
      key: '/dashboard',
      icon: <DashboardOutlined />,
      label: <Link to="/dashboard">Dashboard</Link>,
    },
    {
      key: '/users',
      icon: <UserOutlined />,
      label: <Link to="/users">Users</Link>,
    },
    {
      key: '/customers',
      icon: <TeamOutlined />,
      label: <Link to="/customers">Customers</Link>,
    },
    {
      key: '/settings',
      icon: <SettingOutlined />,
      label: <Link to="/settings">Settings</Link>
    }
  ]