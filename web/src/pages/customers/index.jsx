import React, { useState } from 'react';
import { 
  Typography, 
  Table, 
  Tag, 
  Space, 
  Button, 
  Input, 
  Row, 
  Col,
  Dropdown,
  Modal,
  Form,
  Select,
  DatePicker,
  Tooltip
} from 'antd';
import { 
  SearchOutlined, 
  PlusOutlined, 
  EditOutlined, 
  DeleteOutlined, 
  MoreOutlined,
  EyeOutlined,
  ExportOutlined
} from '@ant-design/icons';

const { Title } = Typography;
const { Option } = Select;

const Customers = () => {
  const [searchText, setSearchText] = useState('');
  const [loading, setLoading] = useState(false);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [modalMode, setModalMode] = useState('add'); // 'add' or 'edit'
  const [form] = Form.useForm();
  const [selectedCustomer, setSelectedCustomer] = useState(null);

  // Sample data
  const data = [
    {
      key: '1',
      id: 'CUS-001',
      name: 'John Smith',
      email: 'john@example.com',
      phone: '+1 (555) 123-4567',
      status: 'active',
      type: 'premium',
      dateJoined: '2023-01-15',
      lastPurchase: '2023-05-20',
      totalSpent: 2450.75
    },
    {
      key: '2',
      id: 'CUS-002',
      name: 'Emma Johnson',
      email: 'emma@example.com',
      phone: '+1 (555) 987-6543',
      status: 'inactive',
      type: 'standard',
      dateJoined: '2023-02-28',
      lastPurchase: '2023-03-10',
      totalSpent: 850.25
    },
    {
      key: '3',
      id: 'CUS-003',
      name: 'Michael Brown',
      email: 'michael@example.com',
      phone: '+1 (555) 456-7890',
      status: 'active',
      type: 'premium',
      dateJoined: '2022-11-05',
      lastPurchase: '2023-06-01',
      totalSpent: 4200.00
    },
    {
      key: '4',
      id: 'CUS-004',
      name: 'Sophia Williams',
      email: 'sophia@example.com',
      phone: '+1 (555) 789-0123',
      status: 'pending',
      type: 'standard',
      dateJoined: '2023-05-10',
      lastPurchase: null,
      totalSpent: 0.00
    },
    {
      key: '5',
      id: 'CUS-005',
      name: 'James Taylor',
      email: 'james@example.com',
      phone: '+1 (555) 234-5678',
      status: 'active',
      type: 'standard',
      dateJoined: '2022-09-15',
      lastPurchase: '2023-05-28',
      totalSpent: 1875.50
    }
  ];

  const handleSearch = (e) => {
    setSearchText(e.target.value);
  };

  const showModal = (mode, record = null) => {
    setModalMode(mode);
    setIsModalVisible(true);
    if (record) {
      setSelectedCustomer(record);
      form.setFieldsValue({
        name: record.name,
        email: record.email,
        phone: record.phone,
        status: record.status,
        type: record.type
      });
    } else {
      form.resetFields();
    }
  };

  const handleCancel = () => {
    setIsModalVisible(false);
    form.resetFields();
  };

  const handleOk = () => {
    form.validateFields()
      .then(values => {
        console.log('Success:', values);
        // Here you would typically save the data
        setIsModalVisible(false);
        form.resetFields();
      })
      .catch(info => {
        console.log('Validate Failed:', info);
      });
  };

  const handleDelete = (record) => {
    Modal.confirm({
      title: 'Are you sure you want to delete this customer?',
      content: `Customer: ${record.name} (${record.id})`,
      okText: 'Yes',
      okType: 'danger',
      cancelText: 'No',
      onOk() {
        console.log('OK');
        // Here you would typically delete the record
      },
    });
  };

  const filteredData = data.filter(
    item =>
      item.name.toLowerCase().includes(searchText.toLowerCase()) ||
      item.email.toLowerCase().includes(searchText.toLowerCase()) ||
      item.id.toLowerCase().includes(searchText.toLowerCase())
  );

  // Action menu for each row
  const actionMenu = (record) => ({
    items: [
      {
        key: '1',
        label: 'View Details',
        icon: <EyeOutlined />,
        onClick: () => console.log('View', record),
      },
      {
        key: '2',
        label: 'Edit',
        icon: <EditOutlined />,
        onClick: () => showModal('edit', record),
      },
      {
        key: '3',
        label: 'Delete',
        icon: <DeleteOutlined />,
        danger: true,
        onClick: () => handleDelete(record),
      },
    ]
  });

  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      sorter: (a, b) => a.id.localeCompare(b.id),
    },
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      sorter: (a, b) => a.name.localeCompare(b.name),
      render: (text, record) => <a onClick={() => console.log('View', record)}>{text}</a>,
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
    },
    {
      title: 'Phone',
      dataIndex: 'phone',
      key: 'phone',
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
      render: (status) => {
        let color = 'green';
        if (status === 'inactive') {
          color = 'volcano';
        } else if (status === 'pending') {
          color = 'geekblue';
        }
        return (
          <Tag color={color}>
            {status.toUpperCase()}
          </Tag>
        );
      },
      filters: [
        { text: 'Active', value: 'active' },
        { text: 'Inactive', value: 'inactive' },
        { text: 'Pending', value: 'pending' },
      ],
      onFilter: (value, record) => record.status.indexOf(value) === 0,
    },
    {
      title: 'Customer Type',
      dataIndex: 'type',
      key: 'type',
      filters: [
        { text: 'Standard', value: 'standard' },
        { text: 'Premium', value: 'premium' },
      ],
      onFilter: (value, record) => record.type.indexOf(value) === 0,
    },
    {
      title: 'Date Joined',
      dataIndex: 'dateJoined',
      key: 'dateJoined',
      sorter: (a, b) => new Date(a.dateJoined) - new Date(b.dateJoined),
    },
    {
      title: 'Last Purchase',
      dataIndex: 'lastPurchase',
      key: 'lastPurchase',
      render: (date) => date || 'N/A',
    },
    {
      title: 'Total Spent',
      dataIndex: 'totalSpent',
      key: 'totalSpent',
      sorter: (a, b) => a.totalSpent - b.totalSpent,
      render: (amount) => `$${amount.toFixed(2)}`,
    },
    {
      title: 'Action',
      key: 'action',
      render: (_, record) => (
        <Dropdown menu={actionMenu(record)} trigger={['click']}>
          <Button type="text" icon={<MoreOutlined />} />
        </Dropdown>
      ),
    },
  ];

  return (
    <div>
      <div className="page-header">
        <Title level={2}>Customers</Title>
        <Space>
          <Tooltip title="Export Data">
            <Button icon={<ExportOutlined />}>Export</Button>
          </Tooltip>
          <Button 
            type="primary" 
            icon={<PlusOutlined />}
            onClick={() => showModal('add')}
          >
            Add Customer
          </Button>
        </Space>
      </div>

      <div className="table-actions">
        <Row gutter={16} style={{ marginBottom: 16 }}>
          <Col span={8}>
            <Input
              placeholder="Search customers"
              prefix={<SearchOutlined />}
              value={searchText}
              onChange={handleSearch}
              allowClear
            />
          </Col>
        </Row>
      </div>

      <Table 
        columns={columns} 
        dataSource={filteredData}
        pagination={{ 
          pageSize: 10,
          showSizeChanger: true,
          showTotal: (total) => `Total ${total} items`
        }}
        loading={loading}
        bordered
        scroll={{ x: 'max-content' }}
      />

      {/* Add/Edit Customer Modal */}
      <Modal
        title={modalMode === 'add' ? 'Add New Customer' : 'Edit Customer'}
        open={isModalVisible}
        onOk={handleOk}
        onCancel={handleCancel}
        width={700}
        destroyOnClose
      >
        <Form
          form={form}
          layout="vertical"
          name="customer_form"
        >
          <Row gutter={16}>
            <Col span={12}>
              <Form.Item
                name="name"
                label="Full Name"
                rules={[{ required: true, message: 'Please enter customer name' }]}
              >
                <Input placeholder="Enter full name" />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item
                name="email"
                label="Email"
                rules={[
                  { required: true, message: 'Please enter email' },
                  { type: 'email', message: 'Please enter a valid email' }
                ]}
              >
                <Input placeholder="Enter email address" />
              </Form.Item>
            </Col>
          </Row>
          <Row gutter={16}>
            <Col span={12}>
              <Form.Item
                name="phone"
                label="Phone Number"
              >
                <Input placeholder="Enter phone number" />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item
                name="status"
                label="Status"
                rules={[{ required: true, message: 'Please select status' }]}
              >
                <Select placeholder="Select status">
                  <Option value="active">Active</Option>
                  <Option value="inactive">Inactive</Option>
                  <Option value="pending">Pending</Option>
                </Select>
              </Form.Item>
            </Col>
          </Row>
          <Row gutter={16}>
            <Col span={12}>
              <Form.Item
                name="type"
                label="Customer Type"
                rules={[{ required: true, message: 'Please select customer type' }]}
              >
                <Select placeholder="Select customer type">
                  <Option value="standard">Standard</Option>
                  <Option value="premium">Premium</Option>
                </Select>
              </Form.Item>
            </Col>
            {modalMode === 'add' && (
              <Col span={12}>
                <Form.Item
                  name="dateJoined"
                  label="Date Joined"
                >
                  <DatePicker style={{ width: '100%' }} />
                </Form.Item>
              </Col>
            )}
          </Row>
        </Form>
      </Modal>
    </div>
  );
};

export default Customers; 